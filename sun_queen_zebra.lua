-- Script to "A Brighter Place"

-- Set up a random number generator
math.randomseed(os.time())

-- A function to generate random numbers between 0 and 1
-- Inputs: max - highest number to return
function randomNum(max)
	return math.random(max)
end

-- A function to generate a random color
function randomColor()
	local r = randomNum(255)
	local g = randomNum(255)
	local b = randomNum(255)
	return {r, g, b}
end

-- A function to set the background color
function setBackground()
	love.graphics.setBackgroundColor(randomColor())
end

-- A function to generate a random position within the screen
function randomPosition()
	local x = randomNum(love.graphics.getWidth())
	local y = randomNum(love.graphics.getHeight())
	return {x, y}
end

-- A function to generate a random shape
function randomShape()
	-- Set a random size
	local size = randomNum(100)

	-- Set a random color
	local color = randomColor()

	-- Generate random coordinates
	local x = randomPosition()[1]
	local y = randomPosition()[2]

	-- Draw the shape
	love.graphics.setColor(color)
	love.graphics.rectangle('fill', x, y, size, size)
end

-- Create the main loop
function love.draw()
	setBackground()
	for i=1, 2000 do
		randomShape()
	end
end