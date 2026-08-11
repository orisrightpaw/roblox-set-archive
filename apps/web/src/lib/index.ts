import { dev } from '$app/env';

export interface User {
    id: number;
    user_name: string;
}

export interface AssetSet {
    id?: number;
    name: string;
    description: string;
    image_asset_id?: number;
    image_asset_updated?: number;
    creator_name: string;
    creator_id?: number;
    subscriber_count?: number;
}

export interface UserResponse extends User {
    owned: AssetSet[];
    subscribed: AssetSet[];
}

export interface PaginatedResult<T> {
    pages: number;
    total: number;
    results: T[];
}

export const API_HOST = dev ? 'http://127.0.0.1:3000' : '';
