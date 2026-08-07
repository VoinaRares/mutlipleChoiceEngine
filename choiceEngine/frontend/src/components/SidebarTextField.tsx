import styles from './Sidebar.module.css'

interface TextFieldProps {
    label: string
    value: string
    onChange: (value: string) => void
}

function SidebarTextField({label, value, onChange}: TextFieldProps) {
    return (
        <label className={styles.field}>
            {label}
            <input value={value} onChange={(event) => onChange(event.target.value)} />
        </label>
    )
}

export default SidebarTextField
