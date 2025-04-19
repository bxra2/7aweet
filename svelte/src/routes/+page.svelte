<script lang="ts">
    import SearchForm from '../components/searchForm.svelte'
    import type { Term } from '$lib/types/term'

    export let data

    let randTerms: Term[] = [...data.randTerms]
</script>

<div class="collumn collumn-4 flex-center">
    <SearchForm />
</div>
<div class="collumn collumn-1 flex-center">
    <p class="headline hl1">جرب البحث عن</p>

    <ul dir="ltr" class="al">
        {#each randTerms as randTerm}
            <li  class="al">
                <a
                    class="suggestion-link"
                    href={`search?q=${randTerm.English}`}
                >
                    {randTerm.English}
                </a>
            </li>
        {/each}
    </ul>
</div>

<style>

</style>
