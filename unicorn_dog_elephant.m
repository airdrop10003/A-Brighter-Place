%This Program creates a 'brighter' image from any given file

%Read in the image
im = imread('input.jpg');
info = imfinfo('input.jpg');

%Display the original image
figure
imshow(im)
title('Original Image')

%Convert to double precision to support operations
im = im2double(im);

%Alter the RGB values
imr = im(:,:,1);
img = im(:,:,2);
imb = im(:,:,3);

%Add the values to the color channels 
imr = imr + 0.25;
img = img + 0.25;
imb = imb + 0.25;

%Display the brighter image
figure
imshow(im)
title('Brighter Image')

%Check RGB values to make sure they are not larger than 1
if max(imr(:)) > 1
    imr = min(im(:), 1);
end
if max(img(:)) > 1
    img = min(im(:), 1);
end
if max(imb(:)) > 1
    imb = min(im(:), 1);
end

%Put the altered RGB values back together
im(:,:,1) = imr;
im(:,:,2) = img;
im(:,:,3) = imb;

%Convert to the original data type
im = im2uint8(im);

%Save the brighter image
imwrite(im, 'brighter_place.jpg', info.Format);

%Display the saved image 
figure
imshow(im)
title('Saved Image')