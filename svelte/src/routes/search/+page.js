// @ts-ignore
export async function load({ fetch, url }) {
    const query = url.searchParams.get('q')
    const incFr = url.searchParams.get('french')
    const incDe = url.searchParams.get('german')
    const incDesc = url.searchParams.get('description')
    const searchParams = new URLSearchParams({
        q: query,
    })
    if (incFr === '1') searchParams.set('french', '1')
    if (incDe === '1') searchParams.set('german', '1')
    if (incDesc === '1') searchParams.set('description', '1')

    const reqURI = `http://localhost:5000/api/search?${searchParams.toString()}`
    const res = await fetch(reqURI)
    const text = await res.text()

    if (!res.ok) {
        return {
            status: res.status,
            error: new Error('Could Not Fetch data...'),
        }
    }

    try {
        const terms = JSON.parse(text)
        return { terms, query }
    } catch (err) {
        return {
            status: 500,
            error: new Error('Invalid JSON response'),
        }
    }
}
