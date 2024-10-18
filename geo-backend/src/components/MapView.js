

import { useEffect, useState, useRef } from 'react';
import { MapContainer, TileLayer, FeatureGroup, GeoJSON } from 'react-leaflet';
import { EditControl } from 'react-leaflet-draw';
import { saveShape, getShapes, uploadFile } from '../api';

function MapView() {
  const [shapes, setShapes] = useState([]);
  const mapRef = useRef();

  // Fetch saved shapes from the backend on component mount
  useEffect(() => {
    const fetchShapes = async () => {
      try {
        const response = await getShapes();
        setShapes(response.data); // Ensure backend returns valid GeoJSON
      } catch (error) {
        console.error('Error fetching shapes', error);
      }
    };
    fetchShapes();
  }, []);

  const onCreated = (e) => {
    const newShape = e.layer.toGeoJSON();
    setShapes([...shapes, newShape]);
  };

  const handleSave = async () => {
    try {
      await saveShape(shapes);
      alert('Shapes saved successfully!');
    } catch (error) {
      console.error('Error saving shapes', error);
      alert('Failed to save shapes.');
    }
  };

  const handleFileUpload = async (e) => {
    const file = e.target.files[0];
    const formData = new FormData();
    formData.append('file', file);

    try {
      await uploadFile(formData);
      alert('File uploaded successfully!');
      window.location.reload(); // Reload to fetch new shapes
    } catch (error) {
      console.error('Error uploading file', error);
      alert('Failed to upload file.');
    }
  };

  return (
    <>
      <div className="map-container">
        <input type="file" onChange={handleFileUpload} className="mb-3" />
        <MapContainer center={[51.505, -0.09]} zoom={13} style={{ height: '80vh' }} ref={mapRef}>
          <TileLayer
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            attribution="&copy; OpenStreetMap contributors"
          />
          <FeatureGroup>
            <EditControl
              position="topright"
              onCreated={onCreated}
              draw={{
                rectangle: true,
                polygon: true,
                circle: true,
                marker: true,
              }}
            />
          </FeatureGroup>
          {shapes.map((shape, index) => (
            <GeoJSON key={index} data={shape} />
          ))}
        </MapContainer>
      </div>
      <br></br>
      <div>
      <button onClick={handleSave} className="btn btn-primary mt-3" >
        Save Shapes
      </button>
      </div>
    </>
  );
}

export default MapView;


