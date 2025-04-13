<script lang="ts">
    import type { Term } from '$lib/types/term.js'
    import TermCard from '../../components/termCard.svelte'
    import SearchForm from '../../components/searchForm.svelte';

    export let data
    // let terms: Term[] = [...data.terms]
    let { terms, query } = data
</script>

<div class="collumns">
    <div class="collumn collumn-2">
        <p class="headline hl9">البحث عن</p>
        <p class="citation">{query}</p>
        <SearchForm/>
    </div>

    <div class="collumn collumn-3">
        {#if terms && terms.length > 0}
            {#each terms as term}
                <TermCard
                    french={term.French}
                    german={term.German}
                    arabic={term.Arabic}
                    english={term.English}
                    desc={term.Description}
                    sourceName={term.Source.NameAr}
                    DomainName={term.Domain.NameAr}
                />
            {/each}
        {:else}
            <p>No posts found.</p>
        {/if}
    </div>
</div>
