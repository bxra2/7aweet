export async function load({ fetch }) {
    const API_BASE = import.meta.env.VITE_API_BASE_URL

    if (!API_BASE) {
        throw new Error('VITE_API_BASE_URL is not defined!')
    }

    const res = await fetch(`${API_BASE}/terms/random`, {
        agent: false,
    })

    if (!res.ok) {
        return {
            status: res.status,
            error: new Error('Could Not Fetch data...'),
        }
    }

    try {
        const randTerms = await res.json()
        return { randTerms }
    } catch (err) {
        return {
            status: 500,
            error: new Error('Invalid JSON response'),
        }
    }
}
