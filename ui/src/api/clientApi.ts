
export default async function clientApi(path: string, method?: string, body?: any) {

    const options: any = { cache: 'no-store' }

    if (method == "post") {
        const formData = new FormData()
        for (const key in body) {
            const value = body[key]
            formData.append(key, value)
        }
        options.method = "POST"
        options.body = formData
    }

    let res: any
    const getParams = { lang: localStorage.getItem('locale') || 'en' }

    try {
        res = await fetch(process.env.NEXT_PUBLIC_API_URL + path + (path.match(/\?/) ? '&' : '?') + new URLSearchParams(getParams), options)
    } catch (e) {
        console.log(e)
    }

    if (res?.ok) {
        return res.json()
    }

    return {}
}
