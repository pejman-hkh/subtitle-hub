"use client"
import { Link } from "@/i18n/navigation";
import { useTranslations } from "next-intl";
import { useSearchParams } from "next/navigation";

export default function Pagination({ pagination, module }: any) {
    const t = useTranslations();
    const params = new URLSearchParams(useSearchParams().toString());
    params.delete("page")
    const qs = params.toString() ? "&" + params.toString() : ""
    
    const totalPages = Math.ceil(pagination?.count / pagination?.size);
    const currentPage = pagination?.page;
    const range = 5;
    const startPage = Math.max(1, currentPage - range);
    const endPage = Math.min(totalPages, currentPage + range);

    return (
        <div className="px-4 bottom-0 right-0 items-center w-full py-4 sm:flex sm:justify-between">
            <div className="flex items-center mb-4 sm:mb-0">
                <Link href={`${module}?page=1${qs}`} className="inline-flex p-2 text-indigo-500 hover:text-indigo-900">{t("First")}</Link>
                <Link href={`${module}?page=${pagination?.prev}${qs}`} className="inline-flex p-2 text-indigo-500 hover:text-indigo-900">{t("Previous")}</Link>
                {[...Array(endPage - startPage + 1)].map((_, idx) => {
                    const page = startPage + idx;
                    return (
                        <Link key={page} href={`${module}?page=${page}${qs}`} className={`mx-1 p-2 ${page === currentPage ? 'text-white bg-indigo-500' : 'text-indigo-500 hover:text-indigo-900'}`}>{page}</Link>
                    );
                })}
                <Link href={`${module}?page=${pagination?.next}${qs}`} className="inline-flex p-2 text-indigo-500 hover:text-indigo-900">{t("Next")}</Link>
                <Link href={`${module}?page=${totalPages}${qs}`} className="inline-flex p-2 text-indigo-500 hover:text-indigo-900">{t("Last")}</Link>
            </div>
            <span className="text-sm font-normal text-indigo-500 dark:text-indigo-400">
                {t("Showing")} <span className="font-semibold text-indigo-900 dark:text-white">{pagination?.from}-{pagination?.to}</span> {t("of")} <span className="font-semibold text-indigo-900 dark:text-white">{pagination?.count}</span>
            </span>
        </div>
    );
}
