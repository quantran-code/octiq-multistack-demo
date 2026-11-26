declare module 'node:http' {
  export interface AddressInfo {
    port: number
  }

  export interface IncomingMessage {
    method?: string
    url?: string
  }

  export interface ServerResponse {
    writeHead(status: number, headers?: Record<string, string | number>): void
    end(body?: string): void
  }

  export interface Server {
    listen(port: number, callback?: () => void): Server
    close(callback?: (err?: Error) => void): void
    address(): AddressInfo | string | null
  }

  export interface ClientResponse extends IncomingMessage {
    statusCode?: number
    headers: Record<string, string | string[] | undefined>
    on(event: 'data' | 'end', cb: (chunk?: unknown) => void): void
  }

  export interface ClientRequest {
    end(body?: string): void
    on(event: 'error', cb: (err: Error) => void): void
  }

  export interface RequestOptions {
    host?: string
    port?: number
    path?: string
    method?: string
    headers?: Record<string, string>
  }

  export function createServer(handler: (req: IncomingMessage, res: ServerResponse) => void): Server

  export function get(url: string, callback: (res: ClientResponse) => void): void
  export function get(
    url: string,
    options: RequestOptions,
    callback: (res: ClientResponse) => void
  ): void

  export function request(options: RequestOptions, callback: (res: ClientResponse) => void): ClientRequest
}
