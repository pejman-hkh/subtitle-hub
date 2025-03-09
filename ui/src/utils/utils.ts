import { getLocale } from "next-intl/server";
import { headers } from "next/headers";

export function uniqid(prefix = "", random = false) {
    const sec = Date.now() * 1000 + Math.random() * 1000;
    const id = sec.toString(16).replace(/\./g, "").padEnd(14, "0");
    return `${prefix}${id}${random ? `.${Math.trunc(Math.random() * 100000000)}` : ""}`;
}

export async function currentUrl() {
    const headersList: any = await headers();
    let path = headersList.get("x-pathname") as string
    path = path.replace(/^(\/(fa|en))/g, "")

    //console.log('paaaaaaaaathhhhhhhhhhh', path)
    const searchParams = new URL(headersList.get("x-url"))?.searchParams.toString()

    const lang = await getLocale()

    if (path.substr(0, 1) !== '/')
        path = '/' + path

    let url = path

    if (url.substr(-1) === '/')
        url = url.substr(0, -1)

    let search = searchParams
    if (search)
        search = '&' + search

    const params = (path.match(/\?/g) ? '&' : '?') + new URLSearchParams({ "lang": lang }) + search
    url = path + params

    return url
}