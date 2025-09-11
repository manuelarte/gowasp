import type { User } from '@/models/users.model.ts'
import axios, { type AxiosInstance } from 'axios'

export interface ApiClient {
  login: (username: string, password: string) => Promise<User>
  logout: () => Promise<void>
}

export class HttpClient implements ApiClient {
  private client: AxiosInstance

  constructor (baseURL: string) {
    this.client = axios.create({
      baseURL,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }

  async login (username: string, password: string): Promise<User> {
    const response = await this.client.post<User>('api/users/login', {
      username,
      password,
    })
    return response.data
  }

  async logout (): Promise<void> {
    const response = await this.client.delete<void>('api/users/logout')
    return response.data
  }
}
