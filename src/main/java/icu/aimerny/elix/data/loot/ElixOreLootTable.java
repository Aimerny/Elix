package icu.aimerny.elix.data.loot;

import icu.aimerny.elix.registry.ModBlock;
import net.fabricmc.fabric.api.datagen.v1.FabricDataOutput;
import net.fabricmc.fabric.api.datagen.v1.provider.FabricBlockLootTableProvider;

public class ElixOreLootTable extends FabricBlockLootTableProvider {

    public ElixOreLootTable(FabricDataOutput dataOutput) {
        super(dataOutput);
    }

    @Override
    public void generate() {
        addDrop(ModBlock.OBSIDIAN_ORE, drops(ModBlock.OBSIDIAN_ORE.asItem()));
    }
}
