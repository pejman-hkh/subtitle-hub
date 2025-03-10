import api from "@/api/api"
import { Metadata } from "next"
import { getTranslations } from "next-intl/server"
import { ClientMovie } from "../../movie/[link]/client"

export async function generateMetadata({ params }: {
    params: Promise<{ code: string }>
}): Promise<Metadata> {
    const t = await getTranslations()
    
    const { code } = await params
    const data = await api('/movies/detail?imdb=' + encodeURIComponent(code))
    const movie = data?.data?.movie

    return {
        title: `${t("Download Subtitle")} ${movie?.name} ${movie?.year}`,
        description: `${t("Download Free Subtitle")} ${movie?.name} ${movie?.year}`,
        keywords: `${t("Download Subtitle") + "," + t("Download Free Subtitle")}`
    }
}

export default async function Movie({ params }: {
    params: Promise<{ code: string }>
}) {

    const { code } = await params
    const data = await api('/movies/detail?imdb=' + encodeURIComponent(code))
    const movie = data?.data?.movie

    return <>
        <ClientMovie movie={movie} />
    </>
}