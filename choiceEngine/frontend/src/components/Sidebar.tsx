import {ReactNode} from 'react'
import styles from './Sidebar.module.css'

interface SidebarProps {
    children: ReactNode
    header: string
}

function Sidebar({children, header}: SidebarProps) {
    return (
        <aside className={styles.sidebar}>
            <header className={styles.header}>
                <h2>{header}</h2>
            </header>
            <section className={styles.fields}>{children}</section>
        </aside>
    )
}

export default Sidebar