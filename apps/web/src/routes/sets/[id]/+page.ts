import { API_HOST, type AssetSet } from '$lib';

export async function load({ params, fetch }) {
    const request = await fetch(`${API_HOST}/api/sets/${params.id}`);
    if (request.status === 429) return { set: null, limited: true };
    else if (!request.ok) return { set: null };

    const body = (await request.json()) as AssetSet;

    return { set: body };
}
