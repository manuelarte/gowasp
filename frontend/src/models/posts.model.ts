export interface Post {
  id: number
  createdAt: EpochTimeStamp
  updatedAt: EpochTimeStamp
  postedAt: EpochTimeStamp
  userId: number
  title: string
  content: string
}
