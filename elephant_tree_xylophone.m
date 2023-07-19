%A Brighter Place

% Creating matrix to store data
dataMatrix = zeros(2000,2);

% Populating dataMatrix with 2000 wave samples
for idx = 1:2000
    dataMatrix(idx, 1) = idx;
    dataMatrix(idx, 2) = cos(idx/10);
end

figure('Name', 'A Brighter Place');

% Plotting data points
plot(dataMatrix(:,1), dataMatrix(:,2));

% Adjusting X-axis limits and ticks
xlim([0 2000]);
set(gca, 'XTick', 0:200:2000);

% Adjusting Y-axis limits and ticks
ylim([-1 1]);
set(gca, 'YTick', -1:0.2:1);

% Axes labels
xlabel('Time (samples)');
ylabel('Amplitude');