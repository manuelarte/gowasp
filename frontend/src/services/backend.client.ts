import type { Page } from '@/models/http.model.ts'
import type { Post } from '@/models/posts.model.ts'
import type { User } from '@/models/users.model.ts'
import axios, { type AxiosInstance } from 'axios'
import { wrapper } from 'axios-cookiejar-support'
import { CookieJar } from 'tough-cookie'

export interface ApiClient {
  login: (username: string, password: string) => Promise<User>
  logout: () => Promise<void>

  getStaticPost: (name: string) => Promise<string>
  getPosts: (page: number) => Promise<Page<Post>>
}

export class HttpClient implements ApiClient {
  private client: AxiosInstance

  constructor (baseURL: string) {
    axios.defaults.withCredentials = true
    const jar = new CookieJar()
    this.client = wrapper(axios.create({
      baseURL,
      jar,
      headers: {
        'Content-Type': 'application/json',
      },
    }))
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

  async getStaticPost (name: string): Promise<string> {
    const response = await this.client.get<string>(`static/posts?name=${name}`)
    return response.data
  }

  async getPosts (page: number): Promise<Page<Post>> {
    const size = 3
    const response = await this.client.get<Page<Post>>(`api/posts?page=${page}&size=${size}`)
    return response.data
  }
}
