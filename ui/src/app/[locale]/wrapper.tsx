"use client"
import { usePathname } from "@/i18n/navigation"
import { useEffect, useState } from "react"
import { useTranslations } from "use-intl"

export function Wrapper({ children }: any) {
    const t = useTranslations()
    const [dir, setDir] = useState<string>("ltr")

    const path = usePathname()

    useEffect(() => {
        setDir(t("dir"))
    }, [path])

    return <div dir={dir}>
        {children}
    </div>
}