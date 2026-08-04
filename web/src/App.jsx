import { useState, useEffect } from 'react'
import KnifeTable from './components/KnifeTable.jsx'
import KnifeForm from './components/KnifeForm.jsx'
import { getAllKnives, getKnifeById, createKnife, updateKnife, deleteKnife } from './api.js'
import './style.css'

function App() {
  const [knives, setKnives] = useState([])
  const [editingKnife, setEditingKnife] = useState(null)
  const [error, setError] = useState(null)

  async function loadKnives() {
    try {
      const data = await getAllKnives()
      setKnives(data || [])
      setError(null)
    } catch (e) {
      setError(e.message)
    }
  }

  useEffect(() => {
    loadKnives()
  }, [])

  async function handleSubmit(data) {
    try {
      if (editingKnife) {
        await updateKnife(editingKnife.id, data)
        setEditingKnife(null)
      } else {
        await createKnife(data)
      }
      await loadKnives()
    } catch (e) {
      setError(e.message)
    }
  }

  async function handleDelete(id) {
    if (!confirm('Удалить нож?')) return
    try {
      await deleteKnife(id)
      await loadKnives()
    } catch (e) {
      setError(e.message)
    }
  }

  async function handleEdit(knife) {
    try {
      const fullKnife = await getKnifeById(knife.id)
      setEditingKnife(fullKnife)
      setError(null)
      window.scrollTo({ top: 0, behavior: 'smooth' })
    } catch (e) {
      setError(e.message)
    }
  }

  function handleCancel() {
    setEditingKnife(null)
  }

  return (
    <div className="container">
      <h1>Каталог ножей</h1>

      {error && <p className="error">{error}</p>}

      <KnifeForm
        onSubmit={handleSubmit}
        editingKnife={editingKnife}
        onCancel={handleCancel}
      />

      <KnifeTable
        knives={knives}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </div>
  )
}

export default App
