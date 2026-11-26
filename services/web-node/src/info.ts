import type { IncomingMessage, ServerResponse } from 'node:http'

export type InfoResponse = {
  service: string
  version: string
  status: 'ok'
}

export const SERVICE_NAME = 'gateway-ts'
export const SERVICE_VERSION = '0.1.0'

export function buildInfo(): InfoResponse {
  return { service: SERVICE_NAME, version: SERVICE_VERSION, status: 'ok' }
}

function sendJson(
  res: ServerResponse,
  status: number,
  payload: unknown,
  headers: Record<string, string> = {}
): void {
  res.writeHead(status, { 'Content-Type': 'application/json', ...headers })
  res.end(JSON.stringify(payload))
}

export function handleInfoRequest(req: IncomingMessage, res: ServerResponse): void {
  const path = (req.url ?? '').split('?')[0]

  if (path !== '/info') {
    sendJson(res, 404, { error: 'not found' })
    return
  }

  if (req.method !== 'GET') {
    sendJson(res, 405, { error: 'method not allowed' }, { Allow: 'GET' })
    return
  }

  sendJson(res, 200, buildInfo())
}
