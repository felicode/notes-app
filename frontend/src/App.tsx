import { useState } from 'react'
import './App.css'

function App() {
  const [response, setResponse] = useState('');

  const fetchNotes = async () => {
    try {
      const res = await fetch('/notes')
      const data = await res.json()
      setResponse(JSON.stringify(data, null, 2))
    } catch (err) {
      setResponse(`Request failed: ${String(err)}`)
    }
  }

  return (
    <div style={{ textAlign: 'center' }}>
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          width: '20%',
          margin: '0 auto',
          paddingTop: '150px',
          gap: '20px'
        }}>
          <button onClick={fetchNotes}>Get notes</button>
          <pre style={{ textAlign: 'left' }}>{response}</pre>
      </div>
    </div>
  )
}

export default App
