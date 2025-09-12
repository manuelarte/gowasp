import type { Page } from '@/models/http.model.ts'
import type { Comment, Post } from '@/models/posts.model.ts'
import type { User } from '@/models/users.model.ts'
import axios, { type AxiosInstance } from 'axios'
import { wrapper } from 'axios-cookiejar-support'
import { CookieJar } from 'tough-cookie'

export interface ApiClient {
  login: (username: string, password: string) => Promise<User>
  logout: () => Promise<void>
  signup: (username: string, password: string) => Promise<User>
  getUser: (id: number) => Promise<User>

  getStaticPost: (name: string) => Promise<string>
  getPost: (id: number) => Promise<Post>
  getPosts: (page: number) => Promise<Page<Post>>
  getPostComments: (postId: number) => Promise<Page<Comment>>
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

  async signup (username: string, password: string): Promise<User> {
    const response = await this.client.post<User>('api/users/signup', {
      username,
      password,
    })
    return response.data
  }

  async getUser (userId: number): Promise<User> {
    const response = await this.client.get<User>(`api/users/${userId}`)
    return response.data
  }

  async getStaticPost (name: string): Promise<string> {
    const response = await this.client.get<string>(`static/posts?name=${name}`)
    return response.data
  }

  async getPost (id: number): Promise<Post> {
    // TODO(manuelarte): create endpoint, this endpoint returns the csrf token
    // const response = await this.client.get<Post>(`api/posts/${id}`)
    // return response.data
    return { title: 'Mock', content: 'Mock', id, postedAt: 1, userId: 1, createdAt: 1, updatedAt: 1 } as Post
  }

  async getPostComments (postId: number): Promise<Page<Comment>> {
    const page = 0
    const size = 10
    const response = await this.client.get<Page<Comment>>(`api/posts/${postId}/comments?page=${page}&size=${size}`)
    return response.data
  }

  async getPosts (page: number): Promise<Page<Post>> {
    const size = 3
    const response = await this.client.get<Page<Post>>(`api/posts?page=${page}&size=${size}`)
    return response.data
  }
}
