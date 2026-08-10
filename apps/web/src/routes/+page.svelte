<script>
    import Heading from '$lib/components/Heading.svelte';
    import Image from '$lib/components/Image.svelte';
    import Link from '$lib/components/Link.svelte';
    import Meta from '$lib/components/Meta.svelte';
    import Paragraph from '$lib/components/Paragraph.svelte';
    import Subheading from '$lib/components/Subheading.svelte';

    const Credits = [
        { id: 70315552, username: 'pizzaboxer', credit: 'Original creator and maintainer of the archive' },
        { id: 23902502, username: 'kinery', credit: 'Scraped 33 million users' },
        { id: 39079161, username: 'hitius', credit: 'Scraped 10.5 million users' },
        { id: 1847959411, username: 'Trout', credit: 'Scraped 1 million users' },
        { id: 2485612194, username: 'Multako', credit: 'Scraped 400K users' }
    ];
</script>

<Meta title="Welcome"></Meta>

<article>
    <Heading>Welcome!</Heading>
    <Paragraph>
        This is a remake of <Link href="https://sets.pizzaboxer.xyz">pizzaboxer's Roblox Set Archive</Link>. You can find the <Link
            href="https://github.com/pizzaboxer/roblox-set-archive">original source here</Link
        >, and the <Link href="https://github.com/orisrightpaw/roblox-set-archive">source for this remake here</Link>. This text
        is a work-in-progress.

        <Subheading>What's Changed?</Subheading>
        <p>
            The backend API has been rewritten in Go, originally from ASP.NET Core 6. Why? I wanted to learn Go.<br />
            The frontend has also been entirely rewritten, now in Svelte. (Distributed in source form, too!)<br />
            <br />
            Some <span class="font-bold">breaking API changes</span> have been introduced. All JSON endpoints now return keys
            formatted in <tt>snake_case</tt>, instead of <tt>camelCase</tt>. Please update your project to reflect this when
            switching to this instance.<br />
            <br />
            Furthermore, the original PostgreSQL database dump cannot be used anymore. You must drop all sequences from the database
            before GORM can talk to it. I have no idea why it doesn't support these sequences, but it was easier to just drop them,
            since the database should be read-only at this point.
        </p>
    </Paragraph>

    <Heading>Archive Download</Heading>
    <Paragraph>
        If you wish to download the entire archive, it can be downloaded as <a
            class="text-blue-400 hover:text-blue-300 transition-colors"
            href="https://mega.nz/file/RLQCwKwY#5mAGC92jyjYCyv_GSTcYMCfDaq9IqKg2EGV81uHjWxs">a PostgreSQL dump</a
        > (~813 megabytes when uncompressed).
    </Paragraph>

    <Heading>Credits</Heading>
    <Paragraph>
        These are the people who helped archive all of this data. Without them, the archive would not be nearly as extensive as
        it ended up being.
    </Paragraph>
    <div class="grid grid-cols-4 gap-4">
        {#each Credits as person}
            <div class="flex gap-3">
                <Image src="https://sets.starfall.wtf/api/users/{person.id}/thumbnail" alt="{person.username}'s avatar"></Image>
                <div class="my-auto w-xl">
                    <a
                        class="font-semibold text-xl text-blue-400 hover:text-blue-300 transition-colors"
                        href="https://www.roblox.com/users/{person.id}/profile">{person.username}</a
                    >
                    <p>{person.credit}</p>
                </div>
            </div>
        {/each}
    </div>
</article>
