export interface PageMetadata {
  page: number
  size: number
  totalCount: number
  totalPages: number
}

export interface Page<T> {
  data: T[]
  _metadata: PageMetadata
}
