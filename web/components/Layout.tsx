import { JSX, ComponentChild } from "preact";
import { Head, HeadProps } from "$fresh/runtime.ts";

type LayoutProps =
    & { title: string; }
    & { children: ComponentChild | ComponentChild[]; };

export function Layout({ title, children }: LayoutProps) {
    return (
        <>
            <Head>
                <title>{title}</title>
                <link rel="stylesheet" href="/styles.css" />
            </Head>
            {children}
        </>
    );
}