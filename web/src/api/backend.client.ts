import type { Page } from '@/models/http.model.ts'
import type { Comment, NewComment, Post } from '@/models/posts.model'
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
  getPostComments: (postId: number) => Promise<[string, Page<Comment>]>
  postPostComment: (postId: number, comment: NewComment) => Promise<Comment>
}

export class HttpClient implements ApiClient {
  private client: AxiosInstance

  constructor (baseURL: string = import.meta.env.VITE_BACKEND_URL) {
    const jar = new CookieJar()
    this.client = wrapper(axios.create({
      baseURL,
      jar,
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json',
        'Accept': '*/*',
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
    const response = await this.client.get<string>(
      `static/posts`,
      { params: { name } },
    )
    return response.data
  }

  async getPost (id: number): Promise<Post> {
    const response = await this.client.get<Post>(`api/posts/${id}`)
    return response.data
  }

  async getPostComments (postId: number): Promise<[string, Page<Comment>]> {
    const page = 0
    const size = 10
    const response = await this.client.get<Page<Comment>>(
      `api/posts/${postId}/comments`,
      { params: { page, size } },
    )
    // TODO(manuelarte): I can't make axios to read Set-Cookie header, so I have to send it in another header
    const csrf = response.headers['x-xsrf-token']
    return [csrf, response.data]
  }

  async getPosts (page: number): Promise<Page<Post>> {
    const size = 3
    const response = await this.client.get<Page<Post>>(
      `api/posts`,
      { params: { page, size } },
    )
    return response.data
  }

  async postPostComment (postId: number, comment: NewComment): Promise<Comment> {
    const response = await this.client.post<Comment>(`api/posts/${postId}/comments`, comment)
    return response.data
  }
}
