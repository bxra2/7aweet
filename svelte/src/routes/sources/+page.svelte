<script lang="ts">
    import type { Source } from '$lib/types/source.js'

    export let data
    let sources: Source[] = [...data.sources]
</script>

<div class="info-section col-center">
    {#if sources && sources.length > 0}
        {#each sources as source}
            <div class="collumn collumn-2">
                <div class="head">
                    <span class="headline hl3">
                        <a class="term-link" href={source.URL}>
                            {source.NameAr}
                        </a>
                    </span>
                    <span class="headline hl4">
                        <a class="term-link" href={source.URL}>
                            {source.Name}
                        </a>
                    </span>
                </div>
                <p>
                    {source.Description}
                </p>
            </div>
        {/each}
    {:else}
            <p>لا توجد مصادر.</p>
    {/if}
</div>
