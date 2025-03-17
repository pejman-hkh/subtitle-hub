import api from "@/api/api"
import { currentUrl } from "@/utils/utils"
import { Metadata } from "next"
import { getTranslations } from "next-intl/server"
import { ClientMovie } from "../client"

export async function generateMetadata(): Promise<Metadata> {
    const t = await getTranslations()
    const data = await api(await currentUrl())
    const movie = data?.data?.movie

    return {
        title: `${t("Download Subtitle")} ${movie?.name} ${movie?.year}`,
        description: `${t("Download Free Subtitle")} ${movie?.name} ${movie?.year}`,
        keywords: `${t("Download Subtitle") + "," + t("Download Free Subtitle")}`
    }
}

export default async function Movie() {

    const data = await api(await currentUrl())
    const movie = data?.data?.movie
    const season = data?.data?.season

    return <>
        <ClientMovie movie={movie} season={season} />
    </>
}