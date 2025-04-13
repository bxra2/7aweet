export async function load({ fetch }) {
    const res = await fetch('http://localhost:5000/api/sources')
    //  const contentType = res.headers.get('content-type');
    const text = await res.text()

    if (!res.ok) {
        return {
            status: res.status,
            error: new Error('Could Not Fetch data...'),
        }
    }

    try {
        const sources = JSON.parse(text)
        return { sources }
    } catch (err) {
        return {
            status: 500,
            error: new Error('Invalid JSON response'),
        }
    }
}
