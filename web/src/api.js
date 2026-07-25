const API_URL = 'http://localhost:8080/api/v1/knives'

export async function getAllKnives() {
  const res = await fetch(API_URL)
  if (!res.ok) throw new Error('Ошибка загрузки ножей')
  return res.json()
}

export async function getKnifeById(id) {
  const res = await fetch(`${API_URL}/${id}`)
  if (!res.ok) throw new Error('Нож не найден')
  return res.json()
}

export async function createKnife(knife) {
  const res = await fetch(API_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(knife)
  })
  if (!res.ok) throw new Error('Ошибка создания ножа')
  return res.json()
}

export async function updateKnife(id, knife) {
  const res = await fetch(`${API_URL}/${id}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(knife)
  })
  if (!res.ok) throw new Error('Ошибка обновления ножа')
  return res.json()
}

export async function deleteKnife(id) {
  const res = await fetch(`${API_URL}/${id}`, {
    method: 'DELETE'
  })
  if (!res.ok) throw new Error('Ошибка удаления ножа')
}
