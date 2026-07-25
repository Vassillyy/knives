function KnifeTable({ knives, onEdit, onDelete }) {
  if (knives.length === 0) {
    return <p className="empty">Ножей пока нет. Добавьте первый!</p>
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
          <tr key={knife.id}>
            <td>{knife.name}</td>
            <td>{knife.brand || '—'}</td>
            <td>{knife.material || '—'}</td>
            <td>{knife.blade_length || '—'}</td>
            <td>{knife.handle || '—'}</td>
            <td>{(knife.price / 100).toFixed(2)} ₽</td>
            <td className="desc">{knife.description || '—'}</td>
            <td className="actions">
              <button className="btn-edit" onClick={() => onEdit(knife)}>Ред.</button>
              <button className="btn-delete" onClick={() => onDelete(knife.id)}>Удал.</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default KnifeTable
