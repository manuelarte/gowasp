export interface Post {
  id: number
  createdAt: EpochTimeStamp
  updatedAt: EpochTimeStamp
  postedAt: EpochTimeStamp
  userId: number
  title: string
  content: string
}

export interface Comment {
  id: number
  createdAt: EpochTimeStamp
  updatedAt: EpochTimeStamp
  postedAt: EpochTimeStamp
  postId: number
  userId: number
  comment: string
}

export interface NewComment {
  comment: string
  userId: number
}
