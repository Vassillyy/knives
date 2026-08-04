const API_URL = 'http://localhost:8080/api/v1/knives'

export async function getAllKnives() {
  const res = await fetch(API_URL)
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка загрузки ножей')
  return json.data
}

export async function getKnifeById(id) {
  const res = await fetch(`${API_URL}/${id}`)
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Нож не найден')
  return json.data
}

export async function createKnife(knife) {
  const res = await fetch(API_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(knife)
  })
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка создания ножа')
  return json
}

export async function updateKnife(id, knife) {
  const res = await fetch(`${API_URL}/${id}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(knife)
  })
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка обновления ножа')
  return json
}

export async function deleteKnife(id) {
  const res = await fetch(`${API_URL}/${id}`, {
    method: 'DELETE'
  })
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка удаления ножа')
  return json
}

export function getPhotoUrl(knifeId, photoId) {
  return `${API_URL}/${knifeId}/photos/${photoId}/file`
}

export async function getKnifePhotos(knifeId) {
  const res = await fetch(`${API_URL}/${knifeId}/photos`)
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка загрузки фотографий')
  return json.data
}

export async function uploadKnifePhoto(knifeId, file) {
  const formData = new FormData()
  formData.append('file', file)
  const res = await fetch(`${API_URL}/${knifeId}/photos`, {
    method: 'POST',
    body: formData
  })
  let json
  try {
    json = await res.json()
  } catch {
    throw new Error(`Ошибка загрузки фотографии (код ${res.status})`)
  }
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка загрузки фотографии')
  return json.data
}

export async function deleteKnifePhoto(knifeId, photoId) {
  const res = await fetch(`${API_URL}/${knifeId}/photos/${photoId}`, {
    method: 'DELETE'
  })
  const json = await res.json()
  if (!res.ok || !json.success) throw new Error(json.error || 'Ошибка удаления фотографии')
  return json
}
