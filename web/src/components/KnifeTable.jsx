import { useState, Fragment } from 'react'
import KnifePhotos from './KnifePhotos.jsx'

function KnifeTable({ knives, onEdit, onDelete }) {
  const [expandedId, setExpandedId] = useState(null)

  if (knives.length === 0) {
    return <p className="empty">Ножей пока нет. Добавьте первый!</p>
  }

  function togglePhotos(id) {
    setExpandedId(expandedId === id ? null : id)
  }

  return (
    <table>
      <thead>
        <tr>
          <th>Название</th>
          <th>Бренд</th>
          <th>Материал</th>
          <th>Длина клинка (мм)</th>
          <th>Рукоять</th>
          <th>Цена</th>
          <th>Описание</th>
          <th>Действия</th>
        </tr>
      </thead>
      <tbody>
        {knives.map((knife) => (
          <Fragment key={knife.id}>
            <tr>
              <td>{knife.name}</td>
              <td>{knife.brand || '—'}</td>
              <td>{knife.material || '—'}</td>
              <td>{knife.blade_length || '—'}</td>
              <td>{knife.handle || '—'}</td>
              <td>{(knife.price / 100).toFixed(2)} ₽</td>
              <td className="desc">{knife.description || '—'}</td>
              <td className="actions">
                <button className="btn-photos" onClick={() => togglePhotos(knife.id)}>
                  {expandedId === knife.id ? 'Скрыть' : 'Фото'}
                </button>
                <button className="btn-edit" onClick={() => onEdit(knife)}>Ред.</button>
                <button className="btn-delete" onClick={() => onDelete(knife.id)}>Удал.</button>
              </td>
            </tr>
            {expandedId === knife.id && (
              <tr>
                <td colSpan={8}>
                  <KnifePhotos knifeId={knife.id} />
                </td>
              </tr>
            )}
          </Fragment>
        ))}
      </tbody>
    </table>
  )
}

export default KnifeTable