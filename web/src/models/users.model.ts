export interface User {
  self: string
  id: number
  username: string
  password: string
}

export function getInitials (user: User): string {
  return user.username.slice(0, 2).toUpperCase()
}
