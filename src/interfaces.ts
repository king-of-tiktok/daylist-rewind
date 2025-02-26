export interface UserRecord {
  accessToken: string;
  active: boolean;
  avatar_url: string;
  collectionId: string;
  collectionName: string;
  created: string;
  display_name: string;
  email: string;
  emailVisibility: boolean;
  error_count: number;
  expiry: string;
  id: string;
  refreshToken: string;
  updated: string;
  username: string;
  verified: boolean;
}

export interface PlaylistRecord {
  collectionId: string;
  collectionName: string;
  created: string;
  description: string;
  hash: string;
  id: string;
  owner: string;
  playlist_id: string;
  title: string;
  updated: string;
}
export interface SongRecord {
  collectionId: string;
  collectionName: string;
  created: string;
  updated: string;
  song_id: string;
  name: string;
  artist: string;
  artist_url: string;
  album: string;
  album_url: string;
  album_cover: string;
  preview_url: string;
  duration: number;
}
