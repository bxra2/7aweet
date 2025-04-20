// @ts-ignore
export async function load({ fetch, url }) {
    const query = url.searchParams.get('q')
    const incFr = url.searchParams.get('french')
    const incDe = url.searchParams.get('german')
    const incDesc = url.searchParams.get('description')
    const pageSize = url.searchParams.get('limit') || '10'
    const pageNumber = url.searchParams.get('page')
    const searchParams = new URLSearchParams({
        q: query
    })
    if (incFr === '1') searchParams.set('french', '1')
    if (incDe === '1') searchParams.set('german', '1')
    if (incDesc === '1') searchParams.set('description', '1')

    searchParams.set('limit', pageSize)
    searchParams.set('page', pageNumber)

    const reqURI = `http://localhost:5000/api/search?${searchParams.toString()}`
    const res = await fetch(reqURI)

    const info = await res.json()
    if (!res.ok) {
        return {
            status: res.status,
            error: new Error('Could Not Fetch info...')
        }
    }
    console.log(pageSize)
    try {
        // const terms = JSON.parse(text)
        return { ...info, query, pageSize }
    } catch (err) {
        return {
            status: 500,
            error: new Error('Invalid JSON response')
        }
    }
}
