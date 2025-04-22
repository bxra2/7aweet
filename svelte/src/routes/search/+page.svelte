<script lang="ts">
    import type { Term } from '$lib/types/term.js'
    import TermCard from '../../components/termCard.svelte'
    import SearchForm from '../../components/searchForm.svelte'
    import { page } from '$app/stores'
    import { toArabicIndic } from '$lib/utils/numerals'
    import Paginator from '../../components/paginator.svelte'
    export let data: {
        count: number
        terms: Term[]
        domains: Array<any>
        sources: Array<any>
        query: string
        pageSize: number
    }
    // Destructure the data prop and make it reactive
    let count: number
    let terms: Term[]
    let domains: Array<any>
    let sources: Array<any>
    let query: string
    let limit: number

    $: {
        count = data.count
        terms = data.terms
        domains = data.domains
        sources = data.sources
        query = data.query
        limit = +data.pageSize
    }
    $: currentPage = Number($page.url.searchParams.get('page')) || 1
</script>

<div class="collumns">
    <div class="collumn collumn-2">
        <p class="headline hl1">البحث عن</p>
        <p class="citation">{query}</p>
        <SearchForm />
        <h2 class="">عدد النتائج : {toArabicIndic(count)}</h2>
        <h3 class="">من مصادر</h3>
        <table class="info-table">
            <thead>
                <tr>
                    <th>العدد</th>
                    <th>المصدر (عربي)</th>
                    <th>المصدر (انجليزي)</th>
                </tr>
            </thead>
            {#each sources as source}
                <tbody>
                    <tr>
                        <td align="center"> {toArabicIndic(source.Cnt)}</td>
                        <td>{source.NameAr}</td>
                        <td align="left">{source.Name}</td>
                    </tr>
                </tbody>
            {/each}
        </table>
        <h3 class="">في مجالات</h3>
        <table class="info-table">
            <thead>
                <tr>
                    <th>العدد</th>
                    <th>المصدر (عربي)</th>
                    <th>المصدر (انجليزي)</th>
                </tr>
            </thead>
            {#each domains as domain}
                <tbody>
                    <tr>
                        <td align="center"> {toArabicIndic(domain.Cnt)}</td>
                        <td>{domain.NameAr}</td>
                        <td align="left">{domain.Name}</td>
                    </tr>
                </tbody>
            {/each}
        </table>
    </div>

    <div class="collumn collumn-3">
        <Paginator {query} {currentPage} {count} {limit} />
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
                    sourceURL={term.URL}
                />
            {/each}
        {:else}
            <p>No posts found.</p>
        {/if}
        {#if count > 3}
            <Paginator {query} {currentPage} {count} {limit} />
        {/if}
    </div>
</div>
