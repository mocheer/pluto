import { fromFile} from 'geotiff';

fromFile('./testdata/20230812_164500压缩.tiff')
  .then(tiff => { 
    const image = tiff.getImage(0).then(async image=>{
     
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
      const data = await image.readRasters({
        samples: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23], // 波段数量，一个波段：[0]，三个波段：[2,1,0]
      });
      console.log(data.length,data[0].length)
      console.log(data[1].filter(e=>e).length)
    })
  });