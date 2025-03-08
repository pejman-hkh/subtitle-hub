
export default async function api(path: string) {
    let res: any

    const getParams = new URLSearchParams()

    try {
        res = await fetch(process.env.NEXT_PUBLIC_API_SERVER_URL + path + (path.match(/\?/) ? '&' : '?') + new URLSearchParams(getParams), { cache: "no-cache" })
    } catch (e) {
        console.log(e)
    }

    if (res?.ok) {
        return res.json()
    }

    return {}
}
