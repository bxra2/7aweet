export async function load({ fetch }) {
    const API_BASE = import.meta.env.VITE_API_BASE_URL

    if (!API_BASE) {
        throw new Error('VITE_API_BASE_URL is not defined!')
    }
    const res = await fetch(`${API_BASE}/about`, {
        agent: false,
    })

    //  const contentType = res.headers.get('content-type');
    const text = await res.text()

    if (!res.ok) {
        return {
            status: res.status,
            error: new Error('Could Not Fetch data...'),
        }
    }

    try {
        const domains = JSON.parse(text)
        return { domains }
    } catch (err) {
        return {
            status: 500,
            error: new Error('Invalid JSON response'),
        }
    }
}
