import { fromFile} from 'geotiff';

fromFile('./testdata/ASTGTMV003_N03E112_dem.tif')
  .then(tiff => { 
    const image = tiff.getImage().then(async image=>{
     
      const width = image.getWidth();
      const height = image.getHeight();
      const tileWidth = image.getTileWidth();
      const tileHeight = image.getTileHeight();
      const samplesPerPixel = image.getSamplesPerPixel();
      
      // when we are actually dealing with geo-data the following methods return
      // meaningful results:
      const origin = image.getOrigin();
      const resolution = image.getResolution();
      const bbox = image.getBoundingBox();
  
      console.log(width,height)
      console.log(tileWidth,tileHeight)
      console.log(samplesPerPixel)
      console.log(origin)
      console.log(resolution)
      console.log(origin)
      const data = await image.readRasters();
      console.log(data[0].length)
      console.log(data[0].filter(e=>e).length)
    })
  });