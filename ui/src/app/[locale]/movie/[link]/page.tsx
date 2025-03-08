import api from "@/api/api"
import { currentUrl } from "@/utils/utils"
import { ClientMovie } from "./client"

export default async function Movie() {
    const data = await api(await currentUrl())
    const movie = data?.data?.movie

    return (<ClientMovie movie={movie} />)
}