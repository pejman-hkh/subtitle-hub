
export function formatBytes(bytes: any, decimals = 2, t: any) {
    if (!+bytes) return '0 ' +t("Bytes")

    const k = 1024
    const dm = decimals < 0 ? 0 : decimals
    const sizes = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB']

    const i = Math.floor(Math.log(bytes) / Math.log(k))

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${t(sizes[i])}`
}

export function tn(t: any, str: string) {
    const match: any = str?.match(/^([0-9]+)\s*(.*)/i)
    return match?.[0] ? match[1] + " " + t(match[2]) : t(str)
}