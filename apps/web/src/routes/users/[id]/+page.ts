import { API_HOST, type UserResponse } from '$lib';

export async function load({ params, fetch }) {
    const request = await fetch(`${API_HOST}/api/users/${params.id}`);
    if (request.status === 429) return { user: null, limited: true };
    else if (!request.ok) return { user: null };

    const body = (await request.json()) as UserResponse;

    return { user: body };
}
