import { useState, useEffect, useRef } from 'react'
import { getKnifePhotos, uploadKnifePhoto, deleteKnifePhoto, getPhotoUrl } from '../api.js'

function KnifePhotos({ knifeId }) {
  const [photos, setPhotos] = useState([])
  const [currentIndex, setCurrentIndex] = useState(0)
  const [error, setError] = useState(null)
  const [uploading, setUploading] = useState(false)
  const fileInputRef = useRef(null)

  async function loadPhotos() {
    try {
      const data = await getKnifePhotos(knifeId)
      setPhotos(data || [])
      setCurrentIndex(0)
      setError(null)
    } catch (e) {
      setError(e.message)
    }
  }

  useEffect(() => {
    loadPhotos()
  }, [knifeId])

  async function handleFileChange(e) {
    const file = e.target.files[0]
    if (!file) return
    setUploading(true)
    try {
      await uploadKnifePhoto(knifeId, file)
      await loadPhotos()
      setError(null)
    } catch (e) {
      setError(e.message)
    } finally {
      setUploading(false)
      if (fileInputRef.current) fileInputRef.current.value = ''
    }
  }

  async function handleDelete(photoId) {
    if (!confirm('Удалить фотографию?')) return
    try {
      await deleteKnifePhoto(knifeId, photoId)
      await loadPhotos()
      setError(null)
    } catch (e) {
      setError(e.message)
    }
  }

  function goPrev() {
    setCurrentIndex((i) => (i - 1 + photos.length) % photos.length)
  }

  function goNext() {
    setCurrentIndex((i) => (i + 1) % photos.length)
  }

  const currentPhoto = photos[currentIndex]

  return (
    <div className="knife-photos">
      {error && <p className="error">{error}</p>}

      {photos.length === 0 ? (
        <p className="empty">Фотографий пока нет</p>
      ) : (
        <div className="photo-slider">
          <div className="slider-main">
            {photos.length > 1 && (
              <button type="button" className="slider-nav slider-prev" onClick={goPrev}>‹</button>
            )}
            <img src={getPhotoUrl(knifeId, currentPhoto.id)} alt={currentPhoto.filename} />
            {photos.length > 1 && (
              <button type="button" className="slider-nav slider-next" onClick={goNext}>›</button>
            )}
            <button type="button" className="btn-delete photo-delete" onClick={() => handleDelete(currentPhoto.id)}>×</button>
            {photos.length > 1 && (
              <span className="slider-count">{currentIndex + 1} / {photos.length}</span>
            )}
          </div>

          {photos.length > 1 && (
            <div className="slider-thumbs">
              {photos.map((photo, i) => (
                <img
                  key={photo.id}
                  src={getPhotoUrl(knifeId, photo.id)}
                  alt={photo.filename}
                  className={i === currentIndex ? 'thumb active' : 'thumb'}
                  onClick={() => setCurrentIndex(i)}
                />
              ))}
            </div>
          )}
        </div>
      )}

      <label className="btn-upload">
        {uploading ? 'Загрузка...' : 'Добавить фото'}
        <input
          ref={fileInputRef}
          type="file"
          accept="image/*"
          onChange={handleFileChange}
          disabled={uploading}
          hidden
        />
      </label>
    </div>
  )
}

export default KnifePhotos