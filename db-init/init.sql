-- Import extention to use UUID
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create Accounts table
CREATE TABLE Accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- Generates a random UUID by default
    email VARCHAR(255) UNIQUE NOT NULL,             -- Email must be unique and not null
    password VARCHAR(255) NOT NULL,                 -- Store password as a string (usually hashed)
    name VARCHAR(255) NOT NULL,                     -- Name cannot be null
    faculty VARCHAR(255) NOT NULL,                  -- Faculty information
    type VARCHAR(50)                                -- Account type (admin, user, etc.)
);

-- Create Users table
CREATE TABLE Users (
    id UUID PRIMARY KEY REFERENCES Accounts(id) ON DELETE CASCADE,  -- Reference the Accounts table
    user_id VARCHAR(50) UNIQUE NOT NULL,                            -- User ID must be unique and not null
    role VARCHAR(50) NOT NULL                                      -- Role of the user (student, teacher, etc.)
);

-- Create Spaces table
CREATE TABLE Spaces (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- Generates a random UUID by default
    name VARCHAR(255) NOT NULL,                     -- Name of the space
    description TEXT,                               -- Description of the space
    location FLOAT,                                 -- Why is this a float? IDK but I'm keeping it
    faculty VARCHAR(255) NOT NULL,                  -- Faculty information
    floor INT,                                      -- Floor number of the space
    building VARCHAR(255) NOT NULL,                 -- Building name
    type VARCHAR(50) NOT NULL,                      -- Type of space (classroom, lab, etc.)
    head_staff VARCHAR(255),                        -- Head staff of the space (Why is this a string?)
    is_available BOOLEAN DEFAULT TRUE               -- Is the space available?
);

-- Create Space_Staff_List table
CREATE TABLE Space_Staff_List (
    space_id UUID REFERENCES Spaces(id) ON DELETE CASCADE,   -- Reference the Spaces table
    staff_id UUID REFERENCES Accounts(id) ON DELETE CASCADE,  -- Reference the Accounts table
    PRIMARY KEY (space_id, staff_id)                         -- Composite primary key to ensure uniqueness
);

-- Create Faculty_Access_List table
CREATE TABLE Faculty_Access_List (
    faculty VARCHAR(255) NOT NULL,                          -- Faculty name
    space_id UUID REFERENCES Spaces(id) ON DELETE CASCADE,  -- Reference the Users table
    PRIMARY KEY (faculty, space_id)                          -- Composite primary key to ensure uniqueness
);

-- Create Rooms table
CREATE TABLE Rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),          -- Generates a random UUID by default
    space_id UUID REFERENCES Spaces(id) ON DELETE CASCADE,  -- Reference the Spaces table
    name VARCHAR(255) NOT NULL,                             -- Name of the room
    description TEXT,                                       -- Description of the room
    room_number VARCHAR(50) NOT NULL,                       -- Room number
    capacity INT,                                           -- Capacity of the room
    min_reserve_capacity INT,                               -- Minimum capacity required to reserve the room
    is_available BOOLEAN DEFAULT TRUE                       -- Is the room available?
);

-- Creare Working_Hour_List table
CREATE TABLE Working_Hour_List (
    space_id UUID REFERENCES Spaces(id) ON DELETE CASCADE,  -- Reference the Spaces table
    day_of_week VARCHAR(50) NOT NULL,                       -- Day of the week
    start_time TIME NOT NULL,                               -- Start time of working hours
    end_time TIME NOT NULL,                                 -- End time of working hours
    PRIMARY KEY (space_id, day_of_week)  -- Composite primary key to ensure uniqueness
);

-- Insert some sample data for testing (delete this in production)
-- Generate by ChatGPT
-- Insert sample data into Accounts table
INSERT INTO Accounts (email, password, name, faculty, type)
VALUES
('alice@example.com', 'hashed_password_1', 'Alice Smith', 'Engineering', 'staff'),
('bob@example.com', 'hashed_password_2', 'Bob Johnson', 'Arts', 'staff'),
('carol@example.com', 'hashed_password_3', 'Carol Williams', 'Science', 'staff'),
('dave@example.com', 'hashed_password_4', 'Dave Brown', 'Engineering', 'user'),
('eve@example.com', 'hashed_password_5', 'Eve Davis', 'Arts', 'user'),
('frank@example.com', 'hashed_password_6', 'Frank Miller', 'Science', 'user'),
('grace@example.com', 'hashed_password_7', 'Grace Lee', 'Engineering', 'user'),
('heidi@example.com', 'hashed_password_8', 'Heidi Clark', 'Arts', 'user'),
('ivan@example.com', 'hashed_password_9', 'Ivan White', 'Science', 'user');

-- Insert sample data into Users table
INSERT INTO Users (id, user_id, role)
VALUES
((SELECT id FROM Accounts WHERE email = 'dave@example.com'), 'user_4', 'student'),
((SELECT id FROM Accounts WHERE email = 'eve@example.com'), 'user_5', 'student'),
((SELECT id FROM Accounts WHERE email = 'frank@example.com'), 'user_6', 'teacher'),
((SELECT id FROM Accounts WHERE email = 'grace@example.com'), 'user_7', 'student'),
((SELECT id FROM Accounts WHERE email = 'heidi@example.com'), 'user_8', 'teacher');

-- Insert sample data into Spaces table
INSERT INTO Spaces (name, description, location, faculty, floor, building, type, head_staff, is_available)
VALUES
('Lecture Hall A', 'Main lecture hall with seating for 100', 1.0, 'Engineering', 1, 'Building A', 'classroom', 'Alice Smith', TRUE),
('Science Lab 1', 'Laboratory for science experiments', 2.0, 'Science', 2, 'Building B', 'lab', 'Carol Williams', TRUE),
('Art Studio', 'Creative space for art classes', 3.0, 'Arts', 1, 'Building C', 'classroom', 'Bob Johnson', TRUE),
('Lecture Hall B', 'Secondary lecture hall with seating for 80', 1.0, 'Engineering', 1, 'Building A', 'classroom', 'Alice Smith', TRUE),
('Science Lab 2', 'Advanced lab with equipment for biology', 2.0, 'Science', 2, 'Building B', 'lab', 'Carol Williams', TRUE);

-- Insert sample data into Space_Staff_List table
INSERT INTO Space_Staff_List (space_id, staff_id)
VALUES
((SELECT id FROM Spaces WHERE name = 'Lecture Hall A'), (SELECT id FROM Accounts WHERE email = 'alice@example.com')),
((SELECT id FROM Spaces WHERE name = 'Science Lab 1'), (SELECT id FROM Accounts WHERE email = 'carol@example.com')),
((SELECT id FROM Spaces WHERE name = 'Art Studio'), (SELECT id FROM Accounts WHERE email = 'bob@example.com')),
((SELECT id FROM Spaces WHERE name = 'Lecture Hall B'), (SELECT id FROM Accounts WHERE email = 'alice@example.com')),
((SELECT id FROM Spaces WHERE name = 'Science Lab 2'), (SELECT id FROM Accounts WHERE email = 'carol@example.com'));

-- Insert sample data into Faculty_Access_List table
INSERT INTO Faculty_Access_List (faculty, space_id)
VALUES
('Engineering', (SELECT id FROM Spaces WHERE name = 'Lecture Hall A')),
('Science', (SELECT id FROM Spaces WHERE name = 'Science Lab 1')),
('Arts', (SELECT id FROM Spaces WHERE name = 'Art Studio')),
('Engineering', (SELECT id FROM Spaces WHERE name = 'Lecture Hall B')),
('Science', (SELECT id FROM Spaces WHERE name = 'Science Lab 2'));

-- Insert sample data into Rooms table
INSERT INTO Rooms (space_id, name, description, room_number, capacity, min_reserve_capacity, is_available)
VALUES
((SELECT id FROM Spaces WHERE name = 'Lecture Hall A'), 'Room 101', 'Main room with projector', '101', 50, 10, TRUE),
((SELECT id FROM Spaces WHERE name = 'Science Lab 1'), 'Room 201', 'Lab with chemistry equipment', '201', 30, 5, TRUE),
((SELECT id FROM Spaces WHERE name = 'Art Studio'), 'Room 301', 'Studio with easels and supplies', '301', 20, 5, TRUE),
((SELECT id FROM Spaces WHERE name = 'Lecture Hall B'), 'Room 102', 'Secondary room with seating', '102', 40, 8, TRUE),
((SELECT id FROM Spaces WHERE name = 'Science Lab 2'), 'Room 202', 'Advanced biology lab', '202', 25, 5, TRUE);

-- Insert sample data into Working_Hour_List table
INSERT INTO Working_Hour_List (space_id, day_of_week, start_time, end_time)
VALUES
((SELECT id FROM Spaces WHERE name = 'Lecture Hall A'), 'Monday', '08:00:00', '17:00:00'),
((SELECT id FROM Spaces WHERE name = 'Science Lab 1'), 'Tuesday', '09:00:00', '16:00:00'),
((SELECT id FROM Spaces WHERE name = 'Art Studio'), 'Wednesday', '10:00:00', '15:00:00'),
((SELECT id FROM Spaces WHERE name = 'Lecture Hall B'), 'Thursday', '08:00:00', '17:00:00'),
((SELECT id FROM Spaces WHERE name = 'Science Lab 2'), 'Friday', '09:00:00', '16:00:00');
-- end ChatGPT generated data