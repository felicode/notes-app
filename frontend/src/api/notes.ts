export interface Note { id: number; title: string; body: string }

export async function listNotes(): Promise<Note[]> {
    const res = await fetch('/notes')
    if (!res.ok) throw new Error(`GET /notes failed: ${res.status}`)
    return res.json()
}

export async function createNote(title: string, body: string): Promise<Note> {
    const res = await fetch('/notes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, body }),
    })
    if (!res.ok) throw new Error(`POST /notes failed: ${res.status}`)
    return res.json()
}

export async function updateNote(id: number, title: string, body: string): Promise<void> {
    const res = await fetch(`/notes/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, body }),
    })
    if (!res.ok) throw new Error(`POST /notes/${id} failed: ${res.status}`)
}

export async function deleteNote(id: number): Promise<void> {
    const res = await fetch(`/notes/${id}`, { method: 'DELETE' })
    if (!res.ok) throw new Error(`DELETE /notes/${id} failed: ${res.status}`)
}