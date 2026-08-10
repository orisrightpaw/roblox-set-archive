import type { UserResponse } from '$lib';

export async function load({ params, fetch }) {
    const request = await fetch(`http://127.0.0.1:3000/api/users/${params.id}`);
    if (!request.ok) return { user: null };

    const body = (await request.json()) as UserResponse;

    return { user: body };
}
