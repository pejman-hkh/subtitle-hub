import api from "@/api/api";
import { currentUrl } from "@/utils/utils";
import ClientIndex from "./client";

export default async function Home() {
  const data = await api(await currentUrl())

  return (
    <ClientIndex data={data} />
  );
}
