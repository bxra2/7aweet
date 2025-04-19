<script lang="ts">
    import type { Term } from '$lib/types/term.js'
    import TermCard from '../../components/termCard.svelte'
    import SearchForm from '../../components/searchForm.svelte'

    export let data
    // let terms: Term[] = [...data.terms]
    let { terms, query } = data

    const SourceCountMap = {}
    const DomainCountMap = {}
    terms.forEach((term: Term) => {
        const key = `${term.Source.Name}|${term.Source.NameAr}`
        if (!SourceCountMap[key]) {
            SourceCountMap[key] = {
                en: term.Source.Name,
                ar: term.Source.NameAr,
                cnt: 1
            }
        } else {
            SourceCountMap[key].cnt += 1
        }
    })
    terms.forEach((term: Term) => {
        const key = `${term.Domain.Name}|${term.Domain.NameAr}`
        if (!DomainCountMap[key]) {
            DomainCountMap[key] = {
                en: term.Domain.Name,
                ar: term.Domain.NameAr,
                cnt: 1
            }
        } else {
            DomainCountMap[key].cnt += 1
        }
    })
    const FoundSources = Object.values(SourceCountMap)
    const FoundDomains = Object.values(DomainCountMap)
</script>

<div class="collumns">
    <div class="collumn collumn-2">
        <p class="headline hl1">البحث عن</p>
        <p class="citation">{query}</p>
        <SearchForm />
        <h2 class="">عدد النتائج : {terms.length}</h2>
        <h3 class="">من مصادر</h3>
        <table>
            <thead>
                <tr>
                    <th>المصدر (عربي)</th>
                    <th>المصدر (انجليزي)</th>
                    <th>العدد</th>
                </tr>
            </thead>
            {#each FoundSources as source}
                <tbody>
                    <tr>
                        <td>{source.ar}</td>
                        <td align="left">{source.en}</td>
                        <td align="center"> {source.cnt}</td>
                    </tr>
                </tbody>
            {/each}
        </table>
        <h3 class="">في مجالات</h3>
        <table>
            <thead>
                <tr>
                    <th>المصدر (عربي)</th>
                    <th>المصدر (انجليزي)</th>
                    <th>العدد</th>
                </tr>
            </thead>
            {#each FoundDomains as domain}
                <tbody>
                    <tr>
                        <td>{domain.ar}</td>
                        <td align="left">{domain.en}</td>
                        <td align="center"> {domain.cnt}</td>
                    </tr>
                </tbody>
            {/each}
        </table>
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
                    sourceURL={term.URL}
                />
            {/each}
        {:else}
            <p>No posts found.</p>
        {/if}
    </div>
</div>
