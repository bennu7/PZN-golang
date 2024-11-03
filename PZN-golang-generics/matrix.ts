enum Access {
    READ = 1,    // 00001
    WRITE = 2,   // 00010
    INSERT = 4,  // 00100
    DELETE = 8,  // 01000
    UPDATE = 16  // 10000
}

type AccessRights = {
    read: boolean;
    write: boolean;
    insert: boolean;
    delete: boolean;
    update: boolean;
};

function getAccessRights(value: number): AccessRights {
    return {
        read: (value & Access.READ) !== 0,
        write: (value & Access.WRITE) !== 0,
        insert: (value & Access.INSERT) !== 0,
        delete: (value & Access.DELETE) !== 0,
        update: (value & Access.UPDATE) !== 0,
    };
}


// Contoh:
let metricss: number[] = [32, 32, 32, 32, 32];
metricss.forEach((value, index) => {
    console.log(`Group ${index + 1}:`, getAccessRights(value));
});

