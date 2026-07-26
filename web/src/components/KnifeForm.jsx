import { useState, useEffect } from 'react'

const emptyForm = {
  name: '',
  description: '',
  price: '',
  material: '',
  blade_length: '',
  handle: '',
  brand: ''
}

function KnifeForm({ onSubmit, editingKnife, onCancel }) {
  const [form, setForm] = useState(emptyForm)

  useEffect(() => {
    if (editingKnife) {
      setForm({
        name: editingKnife.name || '',
        description: editingKnife.description || '',
        price: editingKnife.price ? String(editingKnife.price) : '',
        material: editingKnife.material || '',
        blade_length: editingKnife.blade_length ? String(editingKnife.blade_length) : '',
        handle: editingKnife.handle || '',
        brand: editingKnife.brand || ''
      })
    } else {
      setForm(emptyForm)
    }
  }, [editingKnife])

  function handleChange(e) {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  function handleSubmit(e) {
    e.preventDefault()
    if (editingKnife) {
      const data = {}
      if (form.name !== (editingKnife.name || '')) data.name = form.name
      if (form.price !== String(editingKnife.price || '')) data.price = parseInt(form.price) || 0
      if (form.description !== (editingKnife.description || '')) data.description = form.description || null
      if (form.material !== (editingKnife.material || '')) data.material = form.material || null
      if (form.blade_length !== String(editingKnife.blade_length || '')) data.blade_length = form.blade_length ? parseFloat(form.blade_length) : null
      if (form.handle !== (editingKnife.handle || '')) data.handle = form.handle || null
      if (form.brand !== (editingKnife.brand || '')) data.brand = form.brand || null
      onSubmit(data)
    } else {
      const data = { name: form.name, price: parseInt(form.price) || 0 }
      if (form.description) data.description = form.description
      if (form.material) data.material = form.material
      if (form.blade_length) data.blade_length = parseFloat(form.blade_length)
      if (form.handle) data.handle = form.handle
      if (form.brand) data.brand = form.brand
      onSubmit(data)
    }
    setForm(emptyForm)
  }

  return (
    <form onSubmit={handleSubmit} className="knife-form">
      <h2>{editingKnife ? 'Редактировать нож' : 'Добавить нож'}</h2>

      <div className="form-grid">
        <label>
          Название *
          <input type="text" name="name" value={form.name} onChange={handleChange} required />
        </label>

        <label>
          Бренд
          <input type="text" name="brand" value={form.brand} onChange={handleChange} />
        </label>

        <label>
          Материал
          <input type="text" name="material" value={form.material} onChange={handleChange} />
        </label>

        <label>
          Длина клинка (мм)
          <input type="number" step="0.1" name="blade_length" value={form.blade_length} onChange={handleChange} />
        </label>

        <label>
          Рукоять
          <input type="text" name="handle" value={form.handle} onChange={handleChange} />
        </label>

        <label>
          Цена (копейки) *
          <input type="number" name="price" value={form.price} onChange={handleChange} required />
        </label>
      </div>

      <label className="full-width">
        Описание
        <textarea name="description" value={form.description} onChange={handleChange} rows={3} />
      </label>

      <div className="form-buttons">
        <button type="submit" className="btn-submit">
          {editingKnife ? 'Сохранить' : 'Добавить'}
        </button>
        {editingKnife && (
          <button type="button" className="btn-cancel" onClick={onCancel}>Отмена</button>
        )}
      </div>
    </form>
  )
}

export default KnifeForm
