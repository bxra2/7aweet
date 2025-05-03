<script lang="ts">
    let searchQuery = ''
    let searchAlign = true
    let includeFrench = false
    let includeGerman = false
    let includeDesc = false

    function changeDirection(event: Event) {
        // Ensure event.target is not null and is an HTMLInputElement
        const target = event.target as HTMLInputElement | null

        if (target && target.value) {
            const arabicRegex = /[\u0600-\u06FF]/
            searchAlign = arabicRegex.test(target.value)
        } else {
            searchAlign = false
        }
    }

    function handleSubmit(event: SubmitEvent) {
        event.preventDefault()
        const queryParams = new URLSearchParams({
            q: searchQuery.trim(),
            limit: '10',
            page: '1'
        })
        if (includeFrench) {
            queryParams.set('french', '1')
        }
        if (includeGerman) {
            queryParams.set('german', '1')
        }
        if (includeDesc) {
            queryParams.set('description', '1')
        }
        // Change location to trigger the load function with new URL
        window.location.href = `/search?${queryParams.toString()}`
    }
</script>

<form class="newspaper-form margin-center" on:submit={handleSubmit} method="get">
    <h1 class="title">ابحث في المعجم</h1>
    <input
        type="search"
        placeholder="ابحث هنا..."
        class="search-input"
        class:text-right={searchAlign}
        class:text-left={!searchAlign}
        bind:value={searchQuery}
        on:input={changeDirection}
        name="q"
    />
    <p>يشمل:</p>
    <div class="checkbox-group">
        <label>
            <input
                type="checkbox"
                name="french"
                value="1"
                bind:checked={includeFrench}
            />
            فرنسي
        </label>
        <label>
            <input
                id="german"
                type="checkbox"
                name="german"
                value="1"
                bind:checked={includeGerman}
            />
            الماني</label
        >
        <label>
            <input
                type="checkbox"
                name="description"
                value="1"
                bind:checked={includeDesc}
            />
            وصف</label
        >
    </div>
    <button type="submit" class="newspaper-button">ابحث</button>
</form>

<style>
</style>
