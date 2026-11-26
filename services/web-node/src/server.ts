import { createServer } from 'node:http'
import type { Server } from 'node:http'
import { handleInfoRequest } from './info'

export function createGatewayServer(): Server {
  return createServer(handleInfoRequest)
}
