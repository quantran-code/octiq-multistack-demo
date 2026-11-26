import { afterAll, beforeAll, describe, expect, it } from 'vitest'
import { get, request } from 'node:http'
import type { ClientResponse, Server } from 'node:http'
import { createGatewayServer } from './server'

type HttpResult = {
  statusCode: number
  headers: Record<string, string | string[] | undefined>
  body: string
}

function readResponse(res: ClientResponse): Promise<HttpResult> {
  return new Promise<HttpResult>((resolve) => {
    const chunks: string[] = []
    res.on('data', (chunk) => chunks.push(String(chunk)))
    res.on('end', () => {
      resolve({
        statusCode: res.statusCode ?? 0,
        headers: res.headers,
        body: chunks.join('')
      })
    })
  })
}

describe('GET /info', () => {
  let server: Server
  let port = 0

  beforeAll(async () => {
    server = createGatewayServer()
    await new Promise<void>((resolve) => {
      server.listen(0, resolve)
    })

    const address = server.address()
    if (address === null || typeof address === 'string') {
      throw new Error('expected the gateway server to listen on a TCP port')
    }
    port = address.port
  })

  afterAll(async () => {
    await new Promise<void>((resolve, reject) => {
      server.close((err) => (err ? reject(err) : resolve()))
    })
  })

  function getPath(path: string): Promise<HttpResult> {
    return new Promise<HttpResult>((resolve) => {
      get(`http://127.0.0.1:${port}${path}`, { headers: { Connection: 'close' } }, (res) => {
        void readResponse(res).then(resolve)
      })
    })
  }

  function postPath(path: string): Promise<HttpResult> {
    return new Promise<HttpResult>((resolve, reject) => {
      const req = request(
        { host: '127.0.0.1', port, path, method: 'POST', headers: { Connection: 'close' } },
        (res) => {
          void readResponse(res).then(resolve)
        }
      )
      req.on('error', reject)
      req.end()
    })
  }

  it('returns the service info payload', async () => {
    const response = await getPath('/info')

    expect(response.statusCode).toBe(200)
    expect(response.headers['content-type']).toBe('application/json')
    expect(JSON.parse(response.body)).toEqual({
      service: 'gateway-ts',
      version: '0.1.0',
      status: 'ok'
    })
  })

  it('rejects non-GET methods', async () => {
    const response = await postPath('/info')

    expect(response.statusCode).toBe(405)
    expect(response.headers.allow).toBe('GET')
    expect(response.body).not.toContain('"service"')
  })

  it('returns 404 for unknown paths', async () => {
    const response = await getPath('/health')

    expect(response.statusCode).toBe(404)
    expect(JSON.parse(response.body)).toEqual({ error: 'not found' })
  })
})
