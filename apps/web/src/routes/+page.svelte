<script>
    import Heading from '$lib/components/Heading.svelte';
    import Meta from '$lib/components/Meta.svelte';
    import Paragraph from '$lib/components/Paragraph.svelte';

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
        This is an archive of every Roblox set created until approximately June 2016, with some until November 2018. All the
        data here was archived between May 22nd 2022 and June 1st 2022, with about 1.1 million sets archived.<br />
        <br />
        The user search only shows users that have owned at least one set, so if your user doesn't show up then you likely did not
        own any sets. It has also only captured usernames during the archive period, so if you've changed your username you may need
        to search for your past usernames.
    </Paragraph>

    <Heading>So, what are sets?</Heading>
    <Paragraph>
        Sets are a long-forgotten feature on Roblox that allowed users to create personal collections of public-domain models
        and decals on the catalog. There was a publicly browsable catalog of them, and people could subscribe to sets made by
        other people. Back then, sets were heavily integrated into the platform, with the system being used to power the Insert
        tool, Stamper tool, and the Studio toolbox at the time.<br />
        <br />
        They were very widely used in the early 2010s from their initial introduction. However over time, Roblox's gradual shift
        away from their identity as a brickbuilding sandbox game meant that the concept of sets became less relevant, and was eventually
        abandoned.<br />
        <br />
        This website aims to document this history, and to provide another accessible look into what Roblox was like in the early
        2010s.
    </Paragraph>

    <Heading>Sets of interest</Heading>
    <Paragraph>
        The sets used for the Insert tool are located under the Roblox account.<br />
        <br />
        The sets used for the Stamper tool are located under the usabilityman, usabilityman2 and happyhomeman accounts. Additionally,
        there is a gametest stamper set under UsableMan.
    </Paragraph>

    <Heading>InsertService functionality</Heading>
    <Paragraph>
        The web handler at http://www.roblox.com/Game/Tools/InsertAsset.ashx was a LuaWebService XML endpoint that powered
        relevant InsertService functions, and is now defunct. A compatible replacement is available here at
        https://sets.starfall.wtf/Game/Tools/InsertAsset.ashx that you can redirect to if you wish to use this archive in older
        clients.<br />
        <br />
        Examples: Base Sets / Roblox's Sets / Assets for the 'Parts' Set<br />
        <br />
        If you're looking for JSON endpoints to use in your own code, see the API that this website operates on.
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
                <img
                    class="rounded-full outline outline-zinc-700 aspect-square h-24"
                    src="https://sets.starfall.wtf/api/users/{person.id}/thumbnail"
                    alt="{person.username}'s avatar"
                />
                <div class="my-auto">
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
