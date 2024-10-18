import { useState } from 'react';
import { uploadFile } from '../api';

function FileUpload() {
  const [file, setFile] = useState(null);

  const handleUpload = async () => {
    const formData = new FormData();
    formData.append('file', file);

    try {
      await uploadFile(formData);
      alert('File uploaded successfully!');
    } catch (error) {
      console.error('Error uploading file', error);
    }
  };

  return (
    <div className="upload-container">
      <h2>Upload GeoJSON/KML File</h2>
      <input type="file" onChange={(e) => setFile(e.target.files[0])} />
      <button onClick={handleUpload} className="btn btn-success">Upload</button>
    </div>
  );
}

export default FileUpload;
